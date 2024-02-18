/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Check images link(s) in your markdown file(s)",
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
			fmt.Print(nomsFichiers)

			// Créer une expression régulière pour extraire les liens
			regex := regexp.MustCompile(`!\[(.*?)\]\(([^)]+)\)`)

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
					fmt.Println(liens)
				}
			}

			if err := scanner.Err(); err != nil {
				return err
			}

		}
		for _, lien := range liens {
			cheminImage := filepath.Join(imgPath, lien)
			if _, err := os.Stat(cheminImage); os.IsNotExist(err) {
				return fmt.Errorf("Le fichier d'image %s référencé dans le fichier Markdown %s n'existe pas dans le répertoire des images %s", lien, path, imgPath)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)
}
