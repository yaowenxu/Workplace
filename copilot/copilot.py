# Path: copilot/copilot.py
# Time: 2022-12-23 13:39:50

import os

# main function, prints the current directory and print hello world;
def main():
    print("Hello World! This is a test message from the copilot package.")
    print("The current working directory is: " + os.getcwd())

# call main function
if __name__ == "__main__":
    main()