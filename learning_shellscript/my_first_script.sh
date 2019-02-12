#Print text "WELCOME TO MY SHELL SCRIPT"
echo "WELCOME TO MY SHELL SCRIPT"
 
#function print text "GOOD BYE~~" and exit program
Goodbye() {
echo "GOOD BYE~~"
exit
}
 
#function Check file exist , if not will create new file
CheckFileExist() {
if [ -e ./mynewfile.txt ]; then
  echo "File mynewfile.txt already exists"
else
  echo "File mynewfile.txt doesn't exist."
  #loop for user's input
  selection=
  until [ "$selection" = "y" ]; do
    echo "Do you want to create file name mynewfile.txt (y/n): "
    read selection
    case $selection in
      "y" ) > ./mynewfile.txt ;;
      "n" ) Goodbye;;
      * ) echo "Please Enter Only y or n "
    esac
  done
fi
}
 
#funtion Insert data to mynewfile.txt line by line
InsertData(){
until [ "$selection" = ":exit()" ]; do
  echo "Enter your data (Enter :exit() to exit the program): "
  read selection
  case $selection in
     ":exit()" ) cat ./mynewfile.txt
	         Goodbye;;
     * ) echo $selection >> ./mynewfile.txt
  esac
done
}
 
#Call function CheckFileExist
CheckFileExist
 
#Call function InsertData
InsertData
