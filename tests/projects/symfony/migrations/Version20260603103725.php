<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20260603103725 extends AbstractMigration
{
    public function getDescription(): string
    {
        return 'Entities';
    }

    public function up(Schema $schema): void
    {
        $this->addSql('CREATE TABLE category (id INT AUTO_INCREMENT NOT NULL, name VARCHAR(255) NOT NULL, PRIMARY KEY (id)) DEFAULT CHARACTER SET utf8mb4');
        $this->addSql('CREATE TABLE snippet (id INT AUTO_INCREMENT NOT NULL, name VARCHAR(255) NOT NULL, code LONGTEXT NOT NULL, author_id INT NOT NULL, INDEX IDX_961C8CD5F675F31B (author_id), PRIMARY KEY (id)) DEFAULT CHARACTER SET utf8mb4');
        $this->addSql('CREATE TABLE snippet_category (snippet_id INT NOT NULL, category_id INT NOT NULL, INDEX IDX_19C9CBA66E34B975 (snippet_id), INDEX IDX_19C9CBA612469DE2 (category_id), PRIMARY KEY (snippet_id, category_id)) DEFAULT CHARACTER SET utf8mb4');
        $this->addSql('CREATE TABLE user (id INT AUTO_INCREMENT NOT NULL, username VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, PRIMARY KEY (id)) DEFAULT CHARACTER SET utf8mb4');
        $this->addSql('ALTER TABLE snippet ADD CONSTRAINT FK_961C8CD5F675F31B FOREIGN KEY (author_id) REFERENCES user (id)');
        $this->addSql('ALTER TABLE snippet_category ADD CONSTRAINT FK_19C9CBA66E34B975 FOREIGN KEY (snippet_id) REFERENCES snippet (id) ON DELETE CASCADE');
        $this->addSql('ALTER TABLE snippet_category ADD CONSTRAINT FK_19C9CBA612469DE2 FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE CASCADE');
    }

    public function down(Schema $schema): void
    {
        $this->addSql('ALTER TABLE snippet DROP FOREIGN KEY FK_961C8CD5F675F31B');
        $this->addSql('ALTER TABLE snippet_category DROP FOREIGN KEY FK_19C9CBA66E34B975');
        $this->addSql('ALTER TABLE snippet_category DROP FOREIGN KEY FK_19C9CBA612469DE2');
        $this->addSql('DROP TABLE category');
        $this->addSql('DROP TABLE snippet');
        $this->addSql('DROP TABLE snippet_category');
        $this->addSql('DROP TABLE user');
    }
}
