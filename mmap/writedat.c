#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
//write data struct into regular file
struct student{
	char name[20];
	short age;
	float score;
	char sex;
};
int main()
{
	struct student stu[5];
  	mode_t mode;
	mode=umask(000); 
	int fd=open("user.dat",O_RDWR|O_CREAT|O_EXCL,0666);
	if(fd==-1){
		printf("open:%m\n");
		umask(mode);
		exit(-1);
	}
	printf("ok\n");
	memset(stu,0,sizeof(stu));	
	int i=0;
	for(;i<5;i++){
		memcpy(stu[i].name,"tom",strlen("tom")+1);
		stu[i].age=i;
		stu[i].score=89.12f;
		stu[i].sex='m';
		write(fd,&stu[i],sizeof(stu[i]));		
	}
	close(fd);
	umask(mode);
	return 0;	
}