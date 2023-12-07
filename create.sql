-- DataBase作成
-- CREATE DATABASE my_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

-- ユニットテスト用
-- CREATE DATABASE my_blog_test CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;


-- =========================================================================================================================================================


-- ログイン
-- mysql -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE

-- テーブル作成
CREATE TABLE article (id INTEGER AUTO_INCREMENT PRIMARY KEY NOT NULL,title VARCHAR(255) NOT NULL,content VARCHAR(255) NOT NULL,created DATE NOT NULL);

-- ユニットテスト用テーブル作成
CREATE TABLE article_test (id INTEGER AUTO_INCREMENT PRIMARY KEY NOT NULL,title VARCHAR(255) NOT NULL,content VARCHAR(255) NOT NULL,created DATE NOT NULL);

-- AUTO_INCREMENT確認
SELECT AUTO_INCREMENT FROM information_schema.tables WHERE TABLE_SCHEMA = 'sample' AND TABLE_NAME = 'article_test';