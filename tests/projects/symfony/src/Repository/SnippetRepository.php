<?php

namespace App\Repository;

use App\Entity\Snippet;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends ServiceEntityRepository<Snippet>
 */
class SnippetRepository extends ServiceEntityRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Snippet::class);
    }

    /**
     * @return Snippet[]
     */
    public function search(string $query): array
    {
        return $this->createQueryBuilder('s')
            ->andWhere('s.name LIKE :query')
            ->setParameter('query', '%'.$query.'%')
            ->setMaxResults(50) // Limit results for performance
            ->getQuery()
            ->getResult();
    }
}
