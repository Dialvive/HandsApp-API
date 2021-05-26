INSERT INTO
  `country` (
    `ID`,
    `name_de`,
    `name_es`,
    `name_en`,
    `name_fr`,
    `name_it`,
    `name_pt`,
    `abbreviation`
  )
VALUES
  (
    1,
    '`Mexiko`',
    '`México`',
    '`Mexico`',
    '`Mexique`',
    '`Messico`',
    '`Mexico`',
    '`MX`'
  )
INSERT INTO
  `spoken_language` (
    `ID`,
    `name_de`,
    `name_es`,
    `name_en`,
    `name_fr`,
    `name_it`,
    `name_pt`,
    `abbreviation`
  )
VALUES
  (
    1,
    '`Español`',
    '`Español`',
    '`Español`',
    '`Español`',
    '`Español`',
    '`Español`',
    '`es`'
  );
INSERT INTO
  `sign_language` (
    `ID`,
    `name_de`,
    `name_es`,
    `name_en`,
    `name_fr`,
    `name_it`,
    `name_pt`,
    `abbreviation`
  )
VALUES
  (
    1,
    '`Mexikanische Gebärdensprache`',
    '`Lengua de Señas Mexicana`',
    '`Mexican Sign Language`',
    '`Langue mexicaine Sign`',
    '`Messicano Sign Language`',
    '`Língua de sinais mexicana`',
    '`LSM`'
  );
INSERT INTO
  `locale` (
    `ID`,
    `country_ID`,
    `spoken_language_ID`,
    `sign_language_ID`
  )
VALUES
  (1, 1, 1, 1);
INSERT INTO
  `phrase_category` (
    `ID`,
    `name_de`,
    `name_es`,
    `name_en`,
    `name_fr`,
    `name_it`,
    `name_pt`
  )
VALUES
  (
    1,
    '`Schöne Grüße`',
    '`Saludos`',
    '`	Greetings`',
    '`Les salutations`',
    '`Saluti`',
    '`Saudações`'
  );
INSERT INTO
  `phrase` (
    `ID`,
    `locale_ID`,
    `phrase_category_ID`,
    `text_de`,
    `text_es`,
    `text_en`,
    `text_fr`,
    `text_it`,
    `text_pt`,
    `context_de`,
    `context_es`,
    `context_en`,
    `context_fr`,
    `context_it`,
    `context_pt`
  )
VALUES
  (
    1,
    1,
    1,
    '`Freut mich, dich kennenzulernen`',
    '`Encantado de conocerte`',
    '`nice to meet you`',
    '`Enchanté de faire votre connaissance`',
    '`piacere di conoscerti`',
    '`prazer em conhecê-la`',
    '`Wird verwendet, wenn Sie zum ersten Mal treffen sich jemand.`',
    '`Utilizado cuando recién conoces a alguien.`',
    '`Used when you first meet someone.`',
    '`Utilisé lors de la première rencontre quelqu un.`',
    '`Utilizzato prima volta che incontro qualcuno.`',
    '`Usado quando você conhecer alguém.`'
  );
INSERT INTO
  `region` (`ID`, `country_ID`, `name`)
VALUES
  (1, 1, 'Ciudad de México'),
  (2, 2, '`Madrid`');
INSERT INTO
  `word_category` (
    `ID`,
    `name_de`,
    `name_es`,
    `name_en`,
    `name_fr`,
    `name_it`,
    `name_pt`
  )
VALUES
  (
    1,
    'Sport',
    'Deporte',
    'Sport',
    'Sport',
    'Sport',
    'Esporte'
  );
INSERT INTO
  `word` (
    `ID`,
    `word_category_ID`,
    `text_de`,
    `text_es`,
    `text_en`,
    `text_fr`,
    `text_it`,
    `text_pt`,
    `context_de`,
    `context_es`,
    `context_en`,
    `context_fr`,
    `context_it`,
    `context_pt`,
    `definition_de`,
    `definition_es`,
    `definition_en`,
    `definition_fr`,
    `definition_it`,
    `definition_pt`
  )
VALUES
  (
    1,
    1,
    'fußbal',
    'fútbol',
    'football',
    'football',
    'calcio',
    'futebol',
    '.',
    '.',
    '.',
    '.',
    '.',
    '.'
  );