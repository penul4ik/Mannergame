#!/bin/bash

# Из задания не совсем ясно нужно читать содержимое файлов или их названия,
# поэтому я написал 2 варианта скрипта. Не стал обе логики объединять в
# один скрипт, потому что становится слишком много неизвестных

######## Проверяем имена #########

# Ты мне очень нравишься_*_{name}_{height}_{income}_{age}
# Не любим Ивана

# Получаем список файлов из указанной директории
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

ls -1 $PATH_VALENTINE | grep -vE '(Иван|иван)' | grep -E 'Ты\s+мне\s+очень\s+нравишься\s+[А-Яа-яЁё]{2,20}\s+[0-9]{3}*\s+[1-9][0-9]{4,}\s+[0-9]{2}\.txt'
