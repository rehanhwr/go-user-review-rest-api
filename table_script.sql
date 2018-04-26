CREATE TABLE user_review(
	id INT AUTO_INCREMENT PRIMARY KEY,
	order_id CHAR(5) NOT NULL,
	product_id CHAR(5) NOT NULL,
	user_id CHAR(5) NOT NULL,
	rating FLOAT NOT NULL, 
	review TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
