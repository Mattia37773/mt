<?php

namespace App\Controller\Snippet;

use App\Repository\SnippetRepository;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

#[Route('/')]
final class IndexSnippetController extends AbstractController
{
    #[Route(name: 'app_snippet_index', methods: [Request::METHOD_GET])]
    public function index(SnippetRepository $snippetRepository): Response
    {
        return $this->render('snippet/index.html.twig');
    }
}
