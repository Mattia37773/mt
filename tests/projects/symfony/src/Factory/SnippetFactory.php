<?php

namespace App\Factory;

use App\Entity\Snippet;
use Zenstruck\Foundry\Persistence\PersistentObjectFactory;

/**
 * @extends PersistentObjectFactory<Snippet>
 */
final class SnippetFactory extends PersistentObjectFactory
{
    public function __construct()
    {
    }

    #[\Override]
    public static function class(): string
    {
        return Snippet::class;
    }

    #[\Override]
    protected function defaults(): array|callable
    {
        $rawHtml = self::faker()->randomHtml();
        $formattedHtml = preg_replace('/(<\/[a-zA-Z0-9]+>)/', "$1\n", $rawHtml);

        return [
            'author' => UserFactory::random(),
            'code' => $formattedHtml,
            'name' => self::faker()->sentence(),
            'categories' => CategoryFactory::randomRange(1, 5),
        ];
    }

    #[\Override]
    protected function initialize(): static
    {
        return $this;
    }
}
