
CREATE TABLE IF NOT EXISTS deputados (
                                      id SERIAL PRIMARY KEY,
                                      nome TEXT NOT NULL,
                                      partido TEXT NOT NULL,
                                      votos BIGINT NOT NULL
);
INSERT INTO deputados(nome,partido,votos)values('Ana Sousa','PT',4788),
                                                      ('Paulo Viera','PDT',4587),
                                                      ('Julia Silva','PSD',1457);