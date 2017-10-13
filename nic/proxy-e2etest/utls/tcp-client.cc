#include "nic/proxy-e2etest/ntls.hpp"

int bytes_recv;
int port;
char* test_data;


pthread_t server_thread;

void *main_server(void*);
int main_tcp_client(void);

int main(int argv, char* argc[]) {

  if (argv != 3) {
    printf("usage: ./tls port test_data_file\n");
    exit(-1);
  }
  port = atoi(argc[1]);
  printf("Connecting to port %i\n", port);
  test_data = argc[2];


  main_tcp_client();
  return 0;
}

int create_socket() {
  int sockfd;
  struct sockaddr_in dest_addr;
  struct sockaddr_in src_addr;

  sockfd = socket(AF_INET, SOCK_STREAM, 0);

  memset(&(src_addr), '\0', sizeof(src_addr));
  src_addr.sin_family=AF_INET;
  src_addr.sin_port=htons(0xbaba);

  inet_pton(AF_INET, "64.1.0.4", &src_addr.sin_addr.s_addr);

  if ( bind(sockfd, (const struct sockaddr*)&src_addr, sizeof(src_addr)) != 0 ) {
    perror("can't bind port");
    exit(-1);
  }


  memset(&(dest_addr), '\0', sizeof(dest_addr));
  dest_addr.sin_family=AF_INET;
  dest_addr.sin_port=htons(port);

  inet_pton(AF_INET, "64.0.0.1", &dest_addr.sin_addr.s_addr);

  if ( connect(sockfd, (struct sockaddr *) &dest_addr,
               sizeof(struct sockaddr_in)) == -1 ) {
    perror("Connect: ");
    exit(-1);
  }

  return sockfd;
}

void test_tcp(int transport_fd)
{
  clock_t start, end;
  double cpu_time_used;

  int filefd;
  int bytes;
  int totalbytes = 0;
  bytes_recv = 0;
  char buf[16384];

  int res = 0;
  int total_recv = 0;

  start = clock();

  filefd = open(test_data, O_RDONLY);
  totalbytes = 0;

  res = 0;
  total_recv = 0;

  do {
    bytes = read(filefd, buf, sizeof(buf));
    totalbytes += bytes;
    if (bytes > 0) {
      send(transport_fd, buf, bytes, 0);
      printf("Sent bytes so far %i\n", totalbytes);
    } else {
      break;
    }
    res = recv(transport_fd, buf, 1, 0);
    total_recv += res;
    if (res < 0) {
      printf("TCP Read error: %i\n", res);
    } else {
      printf("Received tcp test data: %i %i\n", res, total_recv);
    }
	
  } while(bytes > 0);

  close(filefd);


  end = clock();
  cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;

  printf("TCP talk time: %.02f\n", cpu_time_used);
}

int main_tcp_client() 
{
  int transport_fd = 0;


  transport_fd = create_socket();

  printf("Connected ! - transport fd %d\n", transport_fd);
  while(1) {
    sleep(5);
  }

  // Start tests
  test_tcp(transport_fd);

  while(1) {
   sleep(5);
  }

  close(transport_fd);

  return(0);
}


