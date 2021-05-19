
exports.up = function(knex) {
    return knex.schema.createTable('phrases', (table) => {
        table.increments();
        table.string('phrase');
    })
};

exports.down = function(knex) {
    return knex.schema.dropTable('phrases');
};
