package ecs

// ComponentID уникальный идентификатор типа компонента
type ComponentID uint8

// World структура ECS мира
type World struct {
	entities     map[uint64]struct{} // Все сущности
	components   map[ComponentID]map[uint64]any
	systems      []*System
	nextEntityID uint64
}

// System описание системы
type System struct {
	RequiredComponents []ComponentID
	UpdateFunc         func([]uint64, map[ComponentID]map[uint64]any)
}

// NewWorld создает новый мир
func NewWorld() *World {
	return &World{
		entities:   make(map[uint64]struct{}),
		components: make(map[ComponentID]map[uint64]any),
	}
}

// NewEntity создает новую сущность
func (w *World) NewEntity() uint64 {
	id := w.nextEntityID
	w.entities[id] = struct{}{}
	w.nextEntityID++
	return id
}

func (w *World) AddSystem(s *System) {
	w.systems = append(w.systems, s)
}

// AddComponent добавляет компонент (без рефлексии)
func (w *World) AddComponent(entityID uint64, componentID ComponentID, component any) {
	if _, exists := w.components[componentID]; !exists {
		w.components[componentID] = make(map[uint64]any)
	}
	w.components[componentID][entityID] = component
}

// Update обновляет все системы
func (w *World) Update() {
	for _, sys := range w.systems {
		var matched []uint64

		// Фильтрация сущностей
		for entity := range w.entities {
			hasAll := true
			for _, cid := range sys.RequiredComponents {
				if _, exists := w.components[cid][entity]; !exists {
					hasAll = false
					break
				}
			}
			if hasAll {
				matched = append(matched, entity)
			}
		}

		// Вызов системы
		sys.UpdateFunc(matched, w.components)
	}
}
