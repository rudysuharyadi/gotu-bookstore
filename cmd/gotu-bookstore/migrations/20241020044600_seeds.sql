-- +goose Up
-- +goose StatementBegin
INSERT INTO books (
        id,
        author,
        title,
        category,
        publisher,
        price,
        isbn,
        language,
        publish_date,
        page
    )
VALUES(
        '0192a860-84a4-74d8-ac0f-6f558a55f078',
        'Sarah Andersen',
        'Sarahs Scribbles',
        'webcomic',
        'Tapas Media',
        18.99,
        '123456789012',
        'english',
        CURRENT_TIMESTAMP,
        150
    ),
    (
        '0192a860-84a4-7549-b32a-3ccf5740b625',
        'Phillip K Dick',
        'Do Androids Dream of Electric Sheep?',
        'fantasy',
        'ABC Books',
        24.59,
        '123456789013',
        'english',
        CURRENT_TIMESTAMP,
        160
    ),
    (
        '0192a860-84a4-7397-ac98-cc9a833c1eb7',
        'Douglas Adams',
        'The Hitchhikers Guide to the Galaxy',
        'novels',
        'Entertaiment Media',
        20.99,
        '123456789014',
        'english',
        CURRENT_TIMESTAMP,
        140
    ),
    (
        '0192a860-84a4-7ce8-8eab-2282384781df',
        'Ray Bradbury',
        'Something Wicked This Way Comes',
        'novels',
        'Tapas Media',
        19.99,
        '123456789015',
        'english',
        CURRENT_TIMESTAMP,
        160
    ),
    (
        '0192a860-84a4-7cb4-9ed7-df35cff2d3a1',
        'Sergio Cobo',
        'A Story of Yesterday',
        'webcomic',
        'Tapas Media',
        13.99,
        '123456789016',
        'english',
        CURRENT_TIMESTAMP,
        150
    ),
    (
        '0192a861-73bf-7be3-b46f-a88952e36424',
        'Seth Grahame',
        'Pride and Prejudice',
        'documentary',
        'ABC Books',
        24.99,
        '123456789017',
        'english',
        CURRENT_TIMESTAMP,
        150
    ),
    (
        '0192a861-73bf-7deb-83ec-c913962fd85e',
        'Mark Tim',
        'Curious Cat',
        'webcomic',
        'Tapas Media',
        13.99,
        '123456789018',
        'english',
        CURRENT_TIMESTAMP,
        150
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd