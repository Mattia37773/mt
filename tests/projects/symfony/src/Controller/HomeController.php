<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\HttpFoundation\Request;

class HomeController extends AbstractController
{
    #[Route('/home', name: 'app_home',  methods: [Request::METHOD_GET])]
    public function redirectToSnippets(): Response
    {
        return $this->redirectToRoute('app_snippet_index');
    }
}
