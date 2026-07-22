<?php

namespace App\DataFixtures;

use App\Factory\CategoryFactory;
use App\Factory\SnippetFactory;
use App\Factory\UserFactory;
use Doctrine\Bundle\FixturesBundle\Fixture;
use Doctrine\Persistence\ObjectManager;

class AppFixtures extends Fixture
{
    public function load(ObjectManager $manager): void
    {
        UserFactory::createMany(20);
        CategoryFactory::createMany(50);
        SnippetFactory::createMany(100);
    }
}
