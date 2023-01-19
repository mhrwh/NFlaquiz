source ./.env

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/Colors.csv' INTO TABLE colors FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n';"

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/Areas.csv' INTO TABLE areas FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n';"

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/Countries.csv' INTO TABLE countries FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n' IGNORE 1 LINES;"

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/Quizzes.csv' INTO TABLE quizzes FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\r\n' IGNORE 1 LINES (@1,@2,@3,@4,@5) SET country_name=@1, country_id=@2, hint1=@3, hint2=@4, hint3=@5;"

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/FlagColors.csv' INTO TABLE flag_colors FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\r\n' IGNORE 1 LINES (@1,@2) SET country_id=@1, color_id=@2;"

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/Users.csv' INTO TABLE users FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n';"

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/Results.csv' INTO TABLE results FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n';"