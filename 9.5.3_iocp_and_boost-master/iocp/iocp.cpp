

#include <WinSock2.h>
#include <MSWSock.h>
#include <cstring>
#include <iostream>

#define BUFFER_SIZE 1024
#define THREAD_COUNT 2
#define START_POST_ACCEPTEX 2
#define PORT 8989

enum class IO_OP_TYPE {
	IO_ACCEPT,  // accept
	IO_SEND,
	IO_RECV,
	IO_CONNECT,
	IO_DISCONNECT,
};

struct ServerParams {
	SOCKET listenSocket;
	HANDLE completionPort;
};

typedef struct OverlappedPerIO {
	OVERLAPPED overlapped;
	SOCKET socket;
	WSABUF wsaBuf;
	IO_OP_TYPE type;
	char buffer[BUFFER_SIZE];
} *LPOverlappedPerIO;

void PostAcceptEx(SOCKET listenSocket) {
	SOCKET sock = WSASocket(AF_INET, SOCK_STREAM, 0, NULL, 0, WSA_FLAG_OVERLAPPED);
	if (sock == INVALID_SOCKET) {
		return;
	}
	OverlappedPerIO* overlp = new OverlappedPerIO;
	if (overlp == nullptr) {
		closesocket(sock);
		return;
	}
	ZeroMemory(overlp, sizeof(OverlappedPerIO));
	overlp->socket = sock;
	overlp->wsaBuf.buf = overlp->buffer;
	overlp->wsaBuf.len = BUFFER_SIZE;
	overlp->type = IO_OP_TYPE::IO_ACCEPT;

	DWORD dwByteRecv = 0;
	while (false == AcceptEx(listenSocket,
		sock,
		overlp->wsaBuf.buf,
		0,
		sizeof(SOCKADDR_IN) + 16,
		sizeof(SOCKADDR_IN) + 16,
		&dwByteRecv,
		(LPOVERLAPPED)overlp)) {
		if (WSAGetLastError() == WSA_IO_PENDING) {
			break;
		}
		std::cout << WSAGetLastError() << std::endl;
	}
}

DWORD WINAPI workerThread(LPVOID lpParam) {
	ServerParams* pms = (ServerParams*)lpParam;
	HANDLE completionPort = pms->completionPort;
	SOCKET listenSocket = pms->listenSocket;

	DWORD bytesTrans;
	ULONG_PTR comletionKey;
	LPOverlappedPerIO overlp;
	int ret;
	while (true) {
		BOOL result = GetQueuedCompletionStatus(
			completionPort,
			&bytesTrans,
			&comletionKey,
			(LPOVERLAPPED*)&overlp,
			INFINITE);
		if (!result) {
			if ((GetLastError() == WAIT_TIMEOUT) || (GetLastError() == ERROR_NETNAME_DELETED)) {
				std::cout << "socket disconnection:" << overlp->socket << std::endl;
				closesocket(overlp->socket);
				delete overlp;
				continue;
			}
			std::cout << "GetQueuedCompletionStatus failed" << std::endl;
			return 0;
		}
		switch (overlp->type) {
		case IO_OP_TYPE::IO_ACCEPT:
		{
			PostAcceptEx(listenSocket);
			std::cout << "happed IO_ACCEPT:" << bytesTrans << std::endl;
			setsockopt(overlp->socket, SOL_SOCKET, SO_UPDATE_ACCEPT_CONTEXT, (char*)&(listenSocket), sizeof(SOCKET));

			ZeroMemory(overlp->buffer, BUFFER_SIZE);
			overlp->type = IO_OP_TYPE::IO_RECV;
			overlp->wsaBuf.buf = overlp->buffer;
			overlp->wsaBuf.len = BUFFER_SIZE;
			CreateIoCompletionPort((HANDLE)overlp->socket, completionPort, NULL, 0);

			DWORD dwRecv = 0, dwFlag = 0;
			ret = WSARecv(overlp->socket, &overlp->wsaBuf, 1, &dwRecv, &dwFlag, &(overlp->overlapped), 0);
			if (ret == SOCKET_ERROR && WSAGetLastError() != WSA_IO_PENDING) {
				std::cout << "WSARecv failed:" << WSAGetLastError() << std::endl;
			}
		}
		break;
		case IO_OP_TYPE::IO_RECV:
		{
			std::cout << "happed IO_RECV:" << bytesTrans << std::endl;
			if (bytesTrans == 0) {
				std::cout << "socket disconnection:" << overlp->socket << std::endl;
				closesocket(overlp->socket);
				delete overlp;
				continue;
			}

			std::cout << "recved data:" << overlp->buffer << std::endl;

			ZeroMemory(&overlp->overlapped, sizeof(OVERLAPPED));
			overlp->type = IO_OP_TYPE::IO_SEND;
			overlp->wsaBuf.buf = "response from server\n";
			overlp->wsaBuf.len = strlen("response from server\n");

			DWORD dwSend = 0;
			ret = WSASend(overlp->socket, &overlp->wsaBuf, 1, &dwSend, 0, &(overlp->overlapped), 0);
			if (ret == SOCKET_ERROR && WSAGetLastError() != WSA_IO_PENDING) {
				std::cout << "WSARecv failed:" << WSAGetLastError() << std::endl;
			}
		}
		break;
		case IO_OP_TYPE::IO_SEND:
		{
			std::cout << "happed IO_SEND:" << bytesTrans << std::endl;
			if (bytesTrans == 0) {
				std::cout << "socket disconnection:" << overlp->socket << std::endl;
				closesocket(overlp->socket);
				delete overlp;
				continue;
			}
		}
		break;
		}
	}
	return 0;
}

// 初始化 回滚式
int InitServer(ServerParams& pms) {
	WSADATA wsaData;
	int ret;

	ret = WSAStartup(MAKEWORD(2, 2), &wsaData);
	if (ret == 0) {
		pms.listenSocket = WSASocket(AF_INET, SOCK_STREAM, 0, NULL, 0, WSA_FLAG_OVERLAPPED);
		if (pms.listenSocket != INVALID_SOCKET) {
			// 绑定地址和端口
			sockaddr_in address;
			address.sin_family = AF_INET;
			address.sin_addr.s_addr = INADDR_ANY;
			address.sin_port = htons(PORT);
			ret = bind(pms.listenSocket, (const sockaddr*) & address, sizeof(address));
			if (ret == 0) {
				ret = listen(pms.listenSocket, SOMAXCONN);
				if (ret == 0) {
					pms.completionPort = CreateIoCompletionPort(INVALID_HANDLE_VALUE, NULL, 0, 0);
					if (pms.completionPort != NULL) {
						if (NULL != CreateIoCompletionPort((HANDLE)pms.listenSocket,
							pms.completionPort,
							NULL,
							0)) {
							return 0;
						}
						CloseHandle(pms.completionPort);
					}
				}
			}
		}
		closesocket(pms.listenSocket);
	}
	WSACleanup();
	if (ret == 0) ret = -1;
	return ret;
}


int main() {
	ServerParams pms;
	int ret;
	ret = InitServer(pms);
	if (ret != 0) {
		std::cout << "InitServer Error" << std::endl;
		return 1;
	}

	for (int i = 0; i < THREAD_COUNT; i++) {
		CreateThread(NULL, 0, workerThread, &pms, 0, NULL);
	}

	for (int i = 0; i < START_POST_ACCEPTEX; i++) {
		PostAcceptEx(pms.listenSocket);
	}

	std::cin.get();

	closesocket(pms.listenSocket);
	CloseHandle(pms.completionPort);
	WSACleanup();

	return 0;
}

