<div align="center">
  <a href="https://git.io/typing-svg"><img src="https://readme-typing-svg.herokuapp.com?font=Tektur&duration=4000&color=0077FF&center=true&width=435&lines=WORKER-POOL" alt="Typing SVG" /></a>
</div>

Этот проект представляет собой реализацию пула воркеров в Go. Реализована возможность распределять задачи между несколькими воркерами, работающими параллельно, а также динамичнского добавления и удаления воркеров.

## Установка и использование

1. Склонируйте репозиторий:
   
   ```sh
   git clone https://github.com/Orleansky/worker-pool.git

2. Перейдите в директорию проекта:
   
   ```sh
    cd worker-pool

3. Запустите main.go

    ```sh
    go run main.go

## Структура проекта
• main.go: Основной пакет, который запускает пул воркеров и отправляет задачи.

• pool/pool.go: Пакет, содержащий реализацию пула и воркеров.
