<?php

namespace App\Factory;

use App\Entity\Category;
use Zenstruck\Foundry\Persistence\PersistentObjectFactory;

/**
 * @extends PersistentObjectFactory<Category>
 */
final class CategoryFactory extends PersistentObjectFactory
{
    private const TECH_STACK = [
        'PHP', 'JavaScript', 'TypeScript', 'Python', 'Go', 'Rust', 'Ruby', 'Java', 'C#',
        'Symfony', 'Laravel', 'React', 'Vue.js', 'Next.js', 'Tailwind CSS', 'Bootstrap',
        'Docker', 'Kubernetes', 'MySQL', 'PostgreSQL', 'Redis', 'MongoDB', 'GraphQL',
        'Git', 'GitHub Actions', 'AWS', 'Vercel', 'Nginx', 'Webpack', 'Vite',
    ];

    #[\Override]
    public static function class(): string
    {
        return Category::class;
    }

    protected function defaults(): array|callable
    {
        return [
            'name' => self::faker()->randomElement(self::TECH_STACK),
        ];
    }

    #[\Override]
    protected function initialize(): static
    {
        return $this;
    }
}
