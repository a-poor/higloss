package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const text1 = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras malesuada rutrum enim, eget dignissim nisl placerat vel. Curabitur ultricies sagittis ornare. Sed nibh lorem, imperdiet a tempor sit amet, blandit vitae leo. Aliquam congue, massa in condimentum luctus, ex justo malesuada purus, a blandit purus enim in lacus. Suspendisse vitae fringilla ante. Cras ac consequat erat. Suspendisse potenti. Sed molestie venenatis urna vitae cursus. Phasellus sed hendrerit ante.`

const text2 = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras malesuada rutrum enim, eget dignissim nisl placerat vel. Curabitur ultricies sagittis ornare. Sed nibh lorem, imperdiet a tempor sit amet, blandit vitae leo.`

const text3 = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras malesuada rutrum enim, eget dignissim nisl placerat vel. Curabitur ultricies sagittis ornare. Sed nibh lorem, imperdiet a tempor sit amet, blandit vitae leo. Aliquam congue, massa in condimentum luctus, ex justo malesuada purus, a blandit purus enim in lacus. Suspendisse vitae fringilla ante. `

func main() {
	fmt.Println("Hello, world!")

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 4).
		Width(22)

	pA := style.Render(text1)
	pB := style.Render(text2)
	pC := style.Render(text3)

	joined := lipgloss.JoinHorizontal(lipgloss.Top, pA, pB, pC)

	fmt.Println(joined)
}
