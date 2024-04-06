#!/usr/bin/env bash

#Check permissions
if [ "$(id -u)" != 0 ]
then
    echo [ERROR] Root permissions required >&2
    exit 1
fi


#Variables
file=/var/users
oldIFS=$IFS

#User create function
create_user() {
#Restore IFS
IFS=$oldIFS
#Restore shell
shell=/sbin/nologin

groupadd $group

#Sudoers
if [ "$group" = it ] || [ "$group" = security ]
then
    if ! grep "%$group" /etc/sudoers
    then
        cp /etc/sudoers{,.bkp}
        echo '%'$group 'ALL=(ALL) ALL' >> /etc/sudoers
    fi
    shell=/bin/bash
elif [ "$user" = admin ]
then
    if ! grep "^$user\ " /etc/sudoers
    then
        cp /etc/sudoers{,.bkp}
        echo $user 'ALL=(ALL) ALL' >> /etc/sudoers
    fi
    shell=/bin/bash
fi

mkdir /home/$group
useradd $user -g $group -b /home/$group -s $shell
}

#Check parameters
if [ -f $file ]
then 
    IFS=$'\n'
    for line in $(cat $file)
    do
        user=$(echo $line | cut -d' ' -f1)
        group=$(echo $line | cut -d' ' -f2)
        echo Username: $user Group: $group
        create_user
    done
elif [ -z $2 ] && [ ! -z $1 ]
then
    user=$1
    read -p "Welcome! Print groupname: " group
    echo Username: $user Group: $group
    create_user
elif [ ! -z $2 ]
then
    user=$1
    group=$2
    echo Username: $user Group: $group
    create_user
else
    echo Welcome!
    select option in "Add user" "Show users" "Exit"
    do case $option in
        "Add user")
            read -p "Print username: " user
            read -p "Print groupname: " group
            create_user ;;
        "Show users")
            echo "Current users in system:"
            cut -d: -f1 /etc/passwd | tr '\n' ', ';;
        "Exit")
            break;;
        *) echo Wrong option ;;
        esac
    done
fi
