package reconciliation

import "sort"

type Record struct {
	Reference   string
	AmountCents int64
	Currency    string
}

type Status string

const (
	Matched         Status = "matched"
	Mismatched      Status = "mismatched"
	MissingExternal Status = "missing_external"
	MissingInternal Status = "missing_internal"
)

type Result struct {
	Reference string
	Status    Status
}

func Reconcile(internal []Record, external []Record) []Result {
	internalMap := make(map[string]Record, len(internal))
	externalMap := make(map[string]Record, len(external))
	refs := make(map[string]bool)

	for _, record := range internal {
		internalMap[record.Reference] = record
		refs[record.Reference] = true
	}

	for _, record := range external {
		externalMap[record.Reference] = record
		refs[record.Reference] = true
	}

	results := make([]Result, 0, len(refs))

	for ref := range refs {
		internalRecord, internalOK := internalMap[ref]
		externalRecord, externalOK := externalMap[ref]

		switch {
		case internalOK && externalOK:
			if internalRecord.AmountCents == externalRecord.AmountCents &&
				internalRecord.Currency == externalRecord.Currency {
				results = append(results, Result{Reference: ref, Status: Matched})
			} else {
				results = append(results, Result{Reference: ref, Status: Mismatched})
			}
		case internalOK:
			results = append(results, Result{Reference: ref, Status: MissingExternal})
		case externalOK:
			results = append(results, Result{Reference: ref, Status: MissingInternal})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Reference < results[j].Reference
	})

	return results
}
