<?php

namespace App\Twig\Components;

use App\Repository\SnippetRepository;
use Symfony\UX\LiveComponent\Attribute\AsLiveComponent;
use Symfony\UX\LiveComponent\Attribute\LiveProp;
use Symfony\UX\LiveComponent\DefaultActionTrait;

#[AsLiveComponent]
final class SnippetsIndex
{
    use DefaultActionTrait;

    #[LiveProp(writable: true, url: true)]
    public string $name = '';

    public function __construct(
        private SnippetRepository $snippetRepository,
    ) {
    }

    public function getResults(): array
    {
        return $this->snippetRepository->search($this->name);
    }
}
