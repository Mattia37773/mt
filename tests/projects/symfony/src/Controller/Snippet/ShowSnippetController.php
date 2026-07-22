<?php

declare(strict_types=1);

namespace App\Controller\Snippet;

use App\Entity\Snippet;
use Symfony\Bridge\Doctrine\Attribute\MapEntity;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class ShowSnippetController extends AbstractController
{
    #[Route('/snippet/{slug}', name: 'app_snippet_show', methods: [Request::METHOD_GET])]
    public function show(
        #[MapEntity(mapping: ['slug' => 'slug'])]
        Snippet $snippet,
    ): Response {
        return $this->render('snippet/show.html.twig', [
            'snippet' => $snippet,
        ]);
    }
}
