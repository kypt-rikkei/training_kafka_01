# training_kafka_01

# zookeeper
    + kafka su dung zookeeper de lam gi?
# Broker
    + broker la 1 kafka server, broker nam giu cac topic, moi message duoc gui den broker se duoc broker nem vao 1 topic
    + nhieu kafka broker la kafka cluster
    + kafka broker co ID (dinh nghia trong service "kafka", trong environment, "KAFKA_BROKER_ID=1")
    + broker co broker leader va broker follower, khi broker leader dead thi cac broker con lai se dam nhan vi tri send message vao cac topic
    + 