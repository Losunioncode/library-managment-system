DROP TABLE IF EXISTS Booklist;
DROP TABLE IF EXISTS Userlist;


CREATE TABLE IF NOT EXISTS Booklist (
    Title      VARCHAR(256) NOT NULL,
    ISBN       VARCHAR(16) UNIQUE PRIMARY KEY,
    Author     VARCHAR(256) NOT NULL,
    Publisher  VARCHAR(256) NOT NULL,
    Stock      INT NOT NULL,
    Available  INT NOT NULL,
    RemoveInfo TEXT ,
    CHECK      (stock >= available AND stock >= 0)
);

CREATE TABLE IF NOT EXISTS Userlist(
    ID       VARCHAR(16)  NOT NULL,
    Name     VARCHAR(256) NOT NULL,
    Password VARCHAR(256) NOT NULL,
    Overdue  INT NOT NULL DEFAULT 0,
    Type     INT NOT NULL,
    PRIMARY KEY (ID)
);
INSERT INTO `Userlist`
VALUES ('root', 'admin', 'root', 0, 0);

CREATE TABLE IF NOT EXISTS Recordlist(
    record_id               INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    book_id                 VARCHAR(16) NOT NULL,
    user_id                 VARCHAR(16) NOT NULL,
    IsReturned BOOLEAN      NOT NULL DEFAULT FALSE,
    borrow_date DATETIME    NOT NULL,
    return_date DATETIME,
    deadline DATETIME       NOT NULL,
    extendtimes             INT NOT NULL,
    FOREIGN KEY (book_id)   REFERENCES Booklist(ISBN),
    FOREIGN KEY (user_id)   REFERENCES Userlist(id),
    CHECK (deadline >= borrow_date)
    )AUTO_INCREMENT=1