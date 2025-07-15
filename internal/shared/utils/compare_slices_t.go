package utils

type Identifiable interface {
  GetID() int64
  IsNew() bool
}

// CompareSlices compares two slices of Identifiable items (T) based on their IDs.
// It assumes:
// - Items in slice2 with ID = 0 are new items.
// - Items in slice2 with ID != 0 are existing items (common/updated).
// It returns:
// - updatedItemsInSlice2: Items present in slice2 with non-zero ID. These are the common items, potentially updated.
// - newItemsInSlice2: Items present in slice2 with ID = 0.
// - removedItemsFromSlice1: Items present in slice1 but NOT in slice2 (based on ID).
func CompareSlices[T Identifiable](slice1 []T, slice2 []T) (updatedItemsInSlice2 []T, newItemsInSlice2 []T, removedItemsFromSlice1 []T) {
	// Map slice1 items by ID for efficient lookup
	slice1Map := make(map[int64]T)
	for _, item := range slice1 {
		slice1Map[item.GetID()] = item
	}

	// Map slice2 items by ID for efficient lookup of existing items
	// Only include items with non-zero IDs in this map
	slice2ExistingMap := make(map[int64]T)
	for _, item := range slice2 {
		if item.GetID() == 0 {
			newItemsInSlice2 = append(newItemsInSlice2, item)
		} else {
			slice2ExistingMap[item.GetID()] = item
			updatedItemsInSlice2 = append(updatedItemsInSlice2, item) // These are the common/updated ones
		}
	}

	// Find removed items from slice1
	for _, item1 := range slice1 {
		if _, found := slice2ExistingMap[item1.GetID()]; !found {
			removedItemsFromSlice1 = append(removedItemsFromSlice1, item1)
		}
	}

	return updatedItemsInSlice2, newItemsInSlice2, removedItemsFromSlice1
}
