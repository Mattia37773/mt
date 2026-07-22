<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20260620211915 extends AbstractMigration
{
 public function getDescription(): string
    {
        return 'Seed tech stack categories';
    }
   private const TECH_STACK = [
        'PHP', 'JavaScript', 'TypeScript', 'Python', 'Go', 'Rust', 'Ruby', 'Java', 'C#',
        'Symfony', 'Laravel', 'React', 'Vue.js', 'Next.js', 'Tailwind CSS', 'Bootstrap',
        'Docker', 'Kubernetes', 'MySQL', 'PostgreSQL', 'Redis', 'MongoDB', 'GraphQL',
        'Git', 'GitHub Actions', 'AWS', 'Vercel', 'Nginx', 'Webpack', 'Vite',
    ];

    public function up(Schema $schema): void
    {
        foreach (self::TECH_STACK as $tech) {
            $this->addSql(
                'INSERT INTO category (name) VALUES (:name)',
                ['name' => $tech]
            );
        }
    }

    public function down(Schema $schema): void
    {
        foreach (self::TECH_STACK as $tech) {
            $this->addSql(
                'DELETE FROM category WHERE name = :name',
                ['name' => $tech]
            );
        }
    }
}
