package booty

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// CalculateShares computes the per-regular-crewmember share and the remainder
// when the captain receives twice the share of a regular crewmember.
//
// Parameters:
//   - total: total pieces of eight to distribute (must be >= 0)
//   - crew: total number of crew including the captain (must be > 0)
//
// Returns:
//   - crewShare: pieces of eight each regular crewmember receives
//   - remainder: pieces of eight left undistributed
//
// The divisor is (crew + 1): captain counts as two shares.
func CalculateShares(total, crew int) (crewShare, remainder int) {
	// We assume validation (crew > 0, total >= 0) is handled by the caller.
	divisor := crew + 1 // captain gets two shares
	return total / divisor, total % divisor
}

// AskInt prompts on writer, reads a line from reader, parses it as an int,
// and ensures the value is at least 1. It returns the parsed value or an error.
func AskInt(r io.Reader, w io.Writer, prompt string) (int, error) {
	if _, err := fmt.Fprintf(w, "%s ", prompt); err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("no input provided")
	}
	text := strings.TrimSpace(scanner.Text())
	if text == "" {
		return 0, fmt.Errorf("no input provided")
	}
	v, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("Sorry, I didn't understand %q. Please enter a whole number.", text)
	}
	if v < 1 {
		return 0, fmt.Errorf("Sorry, %q is not allowed. Please enter a number 1 or greater.", text)
	}
	return v, nil
}