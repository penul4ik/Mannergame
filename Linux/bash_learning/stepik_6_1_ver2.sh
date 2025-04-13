#!/bin/bash

# Из задания не совсем ясно нужно читать содержимое файлов или их названия,
# поэтому я написал 2 варианта скрипта. Не стал обе логики объединять в
# один скрипт, потому что становится слишком много неизвестных

######## Проверяем имена #########

# Ты мне очень нравишься_*_{name}_{height}_{income}_{age}
# Не любим Ивана

# Получаем список файлов из указанной директории
# Открываем каждый файл по очереди
# Читаем содержимое файла, выводим путь к файлу - текст валентинки
# Проверяем соответствует ли имя файла заданному шаблону и {name}!=Иван
# Выводим все файлы которые подошли
PATH_VALENTINE=$1

if [ ! -d $PATH_VALENTINE ]; then
	echo "$PATH_VALENTINE not a directory"
	exit 1
elif [ -z $PATH_VALENTINE ]; then
	PATH_VALENTINE="$(pwd)"
fi

#echo "Search valentinki in $PATH_VALENTINE"


for file in $(ls -1 $PATH_VALENTINE | grep -E '*.txt'); do
	valentinka=$(cat $file)
	if [[ "$valentinka" =~ Ты[[:space:]]мне[[:space:]]очень[[:space:]]нравишься[[:space:]][А-Яа-яЁё]{2,20}[[:space:]][0-9]{3}*[[:space:]][1-9][0-9]{4,}[[:space:]][0-9]{2} && ! "$valentinka" =~ .*Иван.* && ! "$valentinka" =~ .*иван.* ]]; then
                	echo "$PATH_VALENTINE$file"
			cat $file
	fi
done
