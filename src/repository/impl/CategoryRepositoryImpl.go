package impl

import (
	"context"
	"errors"
	"log/slog"

	"cloud.google.com/go/firestore"
	"github.com/ruiborda/ecommerce-product-service/src/database"
	"github.com/ruiborda/ecommerce-product-service/src/model"
	"github.com/ruiborda/ecommerce-product-service/src/repository"
	"google.golang.org/api/iterator"
)

// CategoryRepositoryImpl implementa la interfaz CategoryRepository
type CategoryRepositoryImpl struct {
	firestoreClient *firestore.Client
	collectionName  string
}

// NewCategoryRepositoryImpl crea una nueva instancia de CategoryRepositoryImpl
func NewCategoryRepositoryImpl() repository.CategoryRepository {
	// Obtener la conexión de Firestore
	firestoreClient := database.GetFirestoreClient()
	
	// Crear y devolver la instancia del repositorio
	return &CategoryRepositoryImpl{
		firestoreClient: firestoreClient,
		collectionName:  "categories",
	}
}

// CreateCategory crea una nueva categoría en la base de datos
func (r *CategoryRepositoryImpl) CreateCategory(category *model.Category) (*model.Category, error) {
	ctx := context.Background()
	
	// Crear referencia al documento con el ID generado
	docRef := r.firestoreClient.Collection(r.collectionName).Doc(category.Id)
	
	// Escribir el documento en Firestore
	_, err := docRef.Set(ctx, category)
	if err != nil {
		slog.Error("Error creating category", "error", err)
		return nil, err
	}
	
	return category, nil
}

// GetCategoryById obtiene una categoría por su ID
func (r *CategoryRepositoryImpl) GetCategoryById(id string) (*model.Category, error) {
	ctx := context.Background()
	
	// Obtener referencia al documento
	docRef := r.firestoreClient.Collection(r.collectionName).Doc(id)
	docSnap, err := docRef.Get(ctx)
	
	if err != nil {
		if docSnap == nil || !docSnap.Exists() {
			return nil, nil // No existe
		}
		slog.Error("Error fetching category", "id", id, "error", err)
		return nil, err
	}
	
	// Mapear documento a modelo
	var category model.Category
	if err := docSnap.DataTo(&category); err != nil {
		slog.Error("Error mapping category data", "id", id, "error", err)
		return nil, err
	}
	
	return &category, nil
}

// UpdateCategory actualiza una categoría existente
func (r *CategoryRepositoryImpl) UpdateCategory(category *model.Category) (*model.Category, error) {
	ctx := context.Background()
	
	// Obtener referencia al documento
	docRef := r.firestoreClient.Collection(r.collectionName).Doc(category.Id)
	
	// Verificar si el documento existe
	docSnap, err := docRef.Get(ctx)
	if err != nil || !docSnap.Exists() {
		return nil, errors.New("category not found")
	}
	
	// Actualizar el documento
	_, err = docRef.Set(ctx, category)
	if err != nil {
		slog.Error("Error updating category", "id", category.Id, "error", err)
		return nil, err
	}
	
	return category, nil
}

// DeleteCategoryById elimina una categoría por su ID
func (r *CategoryRepositoryImpl) DeleteCategoryById(id string) error {
	ctx := context.Background()
	
	// Obtener referencia al documento
	docRef := r.firestoreClient.Collection(r.collectionName).Doc(id)
	
	// Verificar si el documento existe
	docSnap, err := docRef.Get(ctx)
	if err != nil || !docSnap.Exists() {
		return errors.New("category not found")
	}
	
	// Eliminar el documento
	_, err = docRef.Delete(ctx)
	if err != nil {
		slog.Error("Error deleting category", "id", id, "error", err)
		return err
	}
	
	return nil
}

// GetCategories obtiene todas las categorías
func (r *CategoryRepositoryImpl) GetCategories() ([]*model.Category, error) {
	ctx := context.Background()
	
	// Crear un slice para almacenar los resultados
	var categories []*model.Category
	
	// Obtener todos los documentos de la colección
	iter := r.firestoreClient.Collection(r.collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			slog.Error("Error iterating categories", "error", err)
			return nil, err
		}
		
		// Mapear documento a modelo
		var category model.Category
		if err := doc.DataTo(&category); err != nil {
			slog.Error("Error mapping category data", "id", doc.Ref.ID, "error", err)
			continue
		}
		
		categories = append(categories, &category)
	}
	
	return categories, nil
}