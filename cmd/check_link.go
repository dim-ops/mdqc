/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Check web link(s) in your markdown file(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fichiers, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		var nomsFichiers []string

		for _, fichier := range fichiers {
			if !fichier.IsDir() && strings.HasSuffix(strings.ToLower(fichier.Name()), ".md") {
				nomsFichiers = append(nomsFichiers, fichier.Name())
			}
		}
		fmt.Print(nomsFichiers)

		var liens []string

		for _, fichier := range nomsFichiers {

			f, err := os.Open(path + "/" + fichier)
			if err != nil {
				return err
			}
			defer f.Close()

			// Créer une expression régulière pour extraire les liens
			regex := regexp.MustCompile(`\[(.*?)\]\((https?://[^\s\)]+)\)`)

			// Lire le fichier ligne par ligne
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				ligne := scanner.Text()

				// Trouver tous les liens dans la ligne
				resultats := regex.FindAllStringSubmatch(ligne, -1)
				for _, resultat := range resultats {
					// Le lien est le deuxième groupe capturé
					lien := resultat[2]
					liens = append(liens, lien)
				}
			}

			if err := scanner.Err(); err != nil {
				return err
			}

		}
		for _, lien := range liens {
			// Effectuer une requête HTTP GET pour vérifier le lien
			_, err := http.Get(lien)
			if err != nil {
				fmt.Printf("Erreur en vérifiant le lien %s: %s\n", lien, err)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
