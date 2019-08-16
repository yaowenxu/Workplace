#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <sys/mman.h>

typedef struct {
	char name[20];
	short age;
	float score;
	char sex;
}student;

int main(){
    student *p, *pend;
    int fd;
    fd=open("user.dat",O_RDWR);
    if(fd==-1){
        fd=open("user.dat",O_RDWR|O_CREAT,0666);
        if(fd==-1){
            printf("open and create file failed!\n");
            exit(-1);
        }
    }
    printf("open ok!\n");
    struct stat st;
    int r=fstat(fd,&st);
    if(r==-1){
        printf("file to get regular file's metadata!\n");
        close(fd);
        exit(-1);
    }
    int len=st.st_size;
    p=mmap(NULL,len,PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);
    if(p==NULL || p==(void*)-1){
        printf("file to mapping regular file to virtual memory!");
    }
    pend=p;
    int i=0;
    while(i<(len/sizeof(student))){
        printf("the %d \n",i);
        printf("name=%s\n",p[i].name);
        printf("age=%d\n",p[i].age);
        printf("score=%f\n",p[i].score);
        printf("sex=%c\n",p[i].sex);
        i++;
    }
    munmap(p,len);
    close(fd);
    return 0;
}