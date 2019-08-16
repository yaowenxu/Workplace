#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
 
typedef struct{
	char name[20];
	short age;
	float score;
	char sex;
}Student;
//read binary data from regular file
int main()
{
	Student stu[5];
	mode_t mode;
	mode=umask(0000);
	int fd=open("user.dat",O_RDWR,0666);
	if(fd==-1){
		printf("open:%m\n");
		umask(mode);
		exit(-1);	
	}
	printf("open ok! can read;\n");
	int i=0;
	for(;i<5;i++){
		read(fd,&stu[i],sizeof(stu[i]));
	}
	close(fd);
	i=0;
	for(;i<5;i++){
		printf("stu[%d].name=%s\n",i,stu[i].name);
		printf("stu[%d].age=%d\n",i,stu[i].age);
		printf("stu[%d].sex=%c\n",i,stu[i].sex);
		printf("stu[%d].score=%f\n",i,stu[i].score);		
	}
	umask(mode);
	return 0;
}