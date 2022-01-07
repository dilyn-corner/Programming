// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
    var newResident *Resident
    newResident = &Resident{Name: name, Age: age, Address: address}
    return newResident
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
    if r.Name == "" {
        return false
    }
    if r.Address["street"] == "" {
        return false
    }
    _, ok := r.Address["street"]
    if !ok {
        return false
    }

    return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
    r.Name = ""
    r.Age = 0
    r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
    var i int
    for _, x := range residents {
        if x.HasRequiredInfo() {
            i++
        }
    }
    return i
}
