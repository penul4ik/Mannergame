#!/bin/bash

ps_arr() {
  local input_arr=("$@") # принимаем все аргументы в виде массива
  local prd_image="${input_arr[-1]}" # вычленяем аргумент с конца массива
  unset 'input_arr[-1]' # удаляем последний аргумент из массива всех аргументов

  echo "Массив ${input_arr[@]}" # у нас получился массив только с нужными элементами без конечного аргумента
  echo '\n'
  echo "arg 1 ${prd_image}" # вот наш аргумент к функции

  input_arr+=("new1" "new2" "new3") # меняем изначальный массив

  printf '%s\n' "${input_arr[@]}" # выдаем массив построчно, чтобы mapfile его принял
}

main() {
  local my_files=("file1" "file2") # создаем массив

  echo "Исходный ${my_files[@]}" 

  mapfile -t result < <(ps_arr "${my_files[@]}" "prd_image") # вызываем функцию с раскрытием массива в виде аргументов
                                                             # и записываем в переменную result
  printf '%s\n' "${result[@]}" # выводим построчно
}

main