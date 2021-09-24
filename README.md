1. Resource
   1. Giải thích về các settings trong file docker-compose.yml: https://www.baeldung.com/ops/docker-compose
   2. Giải thích về KAFKA_LISTENERS VÀ KAFKA_ADVERTISED_LISTENERS: https://www.confluent.io/blog/kafka-listeners-explained/
   3. https://blogs.sap.com/2020/11/03/kafka-service-how-it-can-be-used/
2. Giải thích các config trong docker-compose.yml
   1. "version": version của Docker, mỗi version có các cách cài đặt hơi khác nhau
   2. "service": các service và cấu hình theo sau đó, ở đây service có zookeeper và kafka depends-on zookeeper đó
   3. "image": image với các config có sẵn trên Docker hub, khai báo image thì Docker local sẽ pull nó về
   4. "container-name": tên của container để phân biệt với các container khác
   5. "ports": ví dụ '9093:9092', 9093 sẽ là port mà client giao tiếp với host(172.16.210.157), 9092 là port của container, host sẽ giao tiếp với docker container qua port 9092
   6. "depends-on": theo sau là 1 service(1 docker container), để chắc chắn service đang được config được start sau khi service kia start
   7. "environment":
      1. "BROKER_ID": ID của broker, có thể có nhiều broker nên mỗi thằng cần 1 ID, mỗi broker, hay mỗi kafka service cũng là 1 kafka server
      2. 

# zookeeper
    + kafka su dung zookeeper de lam gi?
# Broker
    + broker la 1 kafka server, broker nam giu cac topic, moi message duoc gui den broker se duoc broker nem vao 1 topic
    + nhieu kafka broker la kafka cluster
    + kafka broker co ID (dinh nghia trong service "kafka", trong environment, "KAFKA_BROKER_ID=1")
    + broker co broker leader va broker follower, khi broker leader dead thi cac broker con lai se dam nhan vi tri send message vao cac topic
    ? lam sao de tao nhieu broker, mo nhieu port localhost va tao them bien broker cho consumer?