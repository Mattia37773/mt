<?php

namespace App\Form;

use App\Entity\Category;
use App\Entity\Snippet;
use Symfony\Bridge\Doctrine\Form\Type\EntityType;
use Symfony\Component\Form\AbstractType;
use Symfony\Component\Form\Extension\Core\Type\TextareaType;
use Symfony\Component\Form\Extension\Core\Type\TextType;
use Symfony\Component\Form\FormBuilderInterface;
use Symfony\Component\OptionsResolver\OptionsResolver;
use Symfony\Component\Validator\Constraints\Length;
use Symfony\Component\Validator\Constraints\NotBlank;

class SnippetType extends AbstractType
{
    public function buildForm(FormBuilderInterface $builder, array $options): void
    {
        $builder
            ->add('name', TextType::class, [
                'label' => 'Snippet Titel',
                'constraints' => [
                    new NotBlank(message: 'Bitte gib einen Namen für das Snippet ein.'),
                    new Length(
                        min: 1,
                        max: 255,
                        minMessage: 'Der Name muss mindestens {{ limit }} Zeichen lang sein.',
                        maxMessage: 'Der Name darf maximal {{ limit }} Zeichen lang sein.'
                    ),
                ],
            ])
            ->add('code', TextareaType::class, [
                'label' => 'Dein Code',
                'attr' => [
                    'rows' => 10,
                    'class' => 'font-mono',
                ],
                'constraints' => [
                    new NotBlank(message: 'Das Code-Feld darf nicht leer sein.'),
                    new Length(
                        min: 10,
                        minMessage: 'Der Code muss mindestens {{ limit }} Zeichen lang sein.'
                    ),
                ],
            ])
            ->add('categories', EntityType::class, [
                'class' => Category::class,
                'choice_label' => 'name',
                'multiple' => true,
                'label' => 'Kategorien',
                'constraints' => [
                    new NotBlank(message: 'Bitte wähle mindestens eine Kategorie aus.'),
                ],
            ])
        ;
    }

    public function configureOptions(OptionsResolver $resolver): void
    {
        $resolver->setDefaults([
            'data_class' => Snippet::class,
        ]);
    }
}
