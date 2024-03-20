 #!bin/bash

 ipServer=$2
 nameUser=$1
 patch=$3

 echo "Start script $0"

 if [ "$#" -ne 3 ];
 then
   echo "Need develop 3 arguments";
   exit 1;
 fi

if [ "$nameUser" == "" ];
then
  echo "No arguments number 1 supplied. Need to develop name";
  exit 1;
fi

if [ ! -f "$patch" ]; then
    echo "File not found!";
    exit 1;
fi


if [ "$ipServer" != "" ]; then

 ip=$(whois "$ipServer" 2>/dev/null  | grep -i "No whois server is known for this kind of object.")

  if [ "$ip" ==   "No whois server is known for this kind of object." ];  then
    echo "The arguments number 2  format does not match";
    exit 1;
    fi

else
    echo "No arguments number 1 supplied";
    exit 1;
 fi


 antifraud=$(ssh "$nameUser"@"$ipServer" ps aux 2>/dev/null | grep -wo [a]ntifraud);

if [ "$antifraud" == "antifraud" ]; then
 	echo "Killing antifraud process";
 	ssh "$nameUser"@"$ipServer" "killall $(ssh "$nameUser"@"$ipServer" ps aux 2>/dev/null | grep -iow [a]ntifraud)";
 	echo "Removed old antifraud "
 	ssh "$nameUser"@"$ipServer" "rm antifraud";
 fi

scp  "$patch" "$nameUser"@"$ipServer":.;

ssh "$nameUser"@"$ipServer" "./antifraud";


