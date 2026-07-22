<?php

namespace App\Controller\Snippet;

use App\Entity\Snippet;
use App\Form\SnippetType;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Security\Http\Attribute\IsGranted;

class CreateSnippetController extends AbstractController
{
    #[IsGranted('ROLE_USER')]
    #[Route('/create', name: 'app_snippet_new', methods: [Request::METHOD_GET, Request::METHOD_POST])]
    public function new(Request $request, EntityManagerInterface $entityManager): Response
    {
        $snippet = new Snippet();
        $form = $this->createForm(SnippetType::class, $snippet);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $snippet->setAuthor($this->getUser());
            $entityManager->persist($snippet);
            $entityManager->flush();

            return $this->redirectToRoute('app_snippet_show', [
                'slug' => $snippet->getSlug(),
            ], Response::HTTP_SEE_OTHER);
        }

        return $this->render('snippet/new.html.twig', [
            'snippet' => $snippet,
            'form' => $form,
        ]);
    }
}
