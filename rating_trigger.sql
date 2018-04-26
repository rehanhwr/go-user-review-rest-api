DELIMITER $$
CREATE TRIGGER ratingTrigger BEFORE INSERT ON user_review
FOR EACH ROW
	BEGIN
		if NEW.rating < 1.0 OR NEW.rating > 5.0 then
			signal sqlstate '45000';
		end if;
	end;
$$