package main

import "fmt"

/*
You have information about n different recipes. You are given a string array recipes and a
2D string array ingredients. The ith recipe has the name recipes[i], and you can create it if you
have all the needed ingredients from ingredients[i]. Ingredients to a recipe may need to be created
from other recipes, i.e., ingredients[i] may contain a string that is in recipes.

You are also given a string array supplies containing all the ingredients that you initially have, and
you have an infinite supply of all of them.

Return a list of all the recipes that you can create. You may return the answer in any order.

Note that two recipes may contain each other in their ingredients.


Example 1:

	Input: recipes = ["bread"], ingredients = [["yeast","flour"]], supplies = ["yeast","flour","corn"]
	Output: ["bread"]
	Explanation:
		We can create "bread" since we have the ingredients "yeast" and "flour".

Example 2:

	Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]
	Output: ["bread","sandwich"]
	Explanation:
		We can create "bread" since we have the ingredients "yeast" and "flour".
		We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".

Example 3:

	Input: recipes = ["bread","sandwich","burger"], ingredients = [["yeast","flour"],["bread","meat"],["sandwich","meat","bread"]], supplies = ["yeast","flour","meat"]
	Output: ["bread","sandwich","burger"]
	Explanation:
		We can create "bread" since we have the ingredients "yeast" and "flour".
		We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
		We can create "burger" since we have the ingredient "meat" and can create the ingredients "bread" and "sandwich".


Constraints:

	n == recipes.length == ingredients.length
	1 <= n <= 100
	1 <= ingredients[i].length, supplies.length <= 100
	1 <= recipes[i].length, ingredients[i][j].length, supplies[k].length <= 10
	recipes[i], ingredients[i][j], and supplies[k] consist only of lowercase English letters.
	All the values of recipes and supplies combined are unique.
	Each ingredients[i] does not contain any duplicate values.
*/

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// Step 1: Initialize the graph and in-degree array
	graph := make(map[string][]string)
	inDegree := make(map[string]int)
	available := make(map[string]bool)

	// Mark initial supplies as available
	for _, supply := range supplies {
		available[supply] = true
	}

	// Step 2: Build the graph and in-degree array
	for i, recipe := range recipes {
		for _, ingredient := range ingredients[i] {
			graph[ingredient] = append(graph[ingredient], recipe)
			inDegree[recipe]++
		}
	}

	// Step 3: Initialize the queue with all supplies
	queue := []string{}
	for _, supply := range supplies {
		queue = append(queue, supply)
	}

	// Result list to store the recipes that can be created
	var result []string

	// Step 4: Process the queue for topological sorting
	for len(queue) > 0 {
		ingredient := queue[0]
		queue = queue[1:]

		// Process all recipes that depend on this ingredient
		for _, recipe := range graph[ingredient] {
			inDegree[recipe]--
			if inDegree[recipe] == 0 {
				// All ingredients are available for this recipe
				queue = append(queue, recipe)
				available[recipe] = true
				result = append(result, recipe)
			}
		}
	}

	return result
}

func main() {
	recipes1 := []string{"bread"}
	ingredients1 := [][]string{{"yeast", "flour"}}
	supplies1 := []string{"yeast", "flour", "corn"}
	fmt.Println(findAllRecipes(recipes1, ingredients1, supplies1)) // Output: ["bread"]

	recipes2 := []string{"bread", "sandwich"}
	ingredients2 := [][]string{{"yeast", "flour"}, {"bread", "meat"}}
	supplies2 := []string{"yeast", "flour", "meat"}
	fmt.Println(findAllRecipes(recipes2, ingredients2, supplies2)) // Output: ["bread", "sandwich"]

	recipes3 := []string{"bread", "sandwich", "burger"}
	ingredients3 := [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}
	supplies3 := []string{"yeast", "flour", "meat"}
	fmt.Println(findAllRecipes(recipes3, ingredients3, supplies3)) // Output: ["bread", "sandwich", "burger"]
}
