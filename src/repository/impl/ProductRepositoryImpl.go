// filepath: /home/rui/ecommerce/ecommerce-product-service/src/repository/impl/ProductRepositoryImpl.go
package impl

import (
	"context"
	"github.com/ruiborda/ecommerce-product-service/src/database"
	"github.com/ruiborda/ecommerce-product-service/src/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type ProductRepositoryImpl struct {
	collectionName string
}

func NewProductRepositoryImpl() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		collectionName: "products",
	}
}

func (p *ProductRepositoryImpl) CreateProduct(product *model.Product) (*model.Product, error) {
	ctx := context.Background()
	firestoreClient := database.GetFirestoreClient()

	// Creamos una referencia a la colección de productos
	collection := firestoreClient.Collection(p.collectionName)

	// Insertamos el documento con el ID generado previamente
	_, err := collection.Doc(product.Id).Set(ctx, product)
	if err != nil {
		slog.Error("Error creating product", "error", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductRepositoryImpl) GetProductById(id string) (*model.Product, error) {
	ctx := context.Background()
	firestoreClient := database.GetFirestoreClient()

	// Obtenemos una referencia al documento del producto
	docRef := firestoreClient.Collection(p.collectionName).Doc(id)

	// Obtenemos el documento
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			slog.Info("Product not found", "id", id)
			return nil, nil
		}
		slog.Error("Error getting product", "error", err)
		return nil, err
	}

	var product model.Product
	err = docSnapshot.DataTo(&product)
	if err != nil {
		slog.Error("Error mapping product data", "error", err)
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepositoryImpl) UpdateProduct(product *model.Product) (*model.Product, error) {
	ctx := context.Background()
	firestoreClient := database.GetFirestoreClient()

	// Actualizamos el documento del producto
	_, err := firestoreClient.Collection(p.collectionName).Doc(product.Id).Set(ctx, product)
	if err != nil {
		slog.Error("Error updating product", "error", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductRepositoryImpl) DeleteProductById(id string) error {
	ctx := context.Background()
	firestoreClient := database.GetFirestoreClient()

	// Eliminamos el documento del producto
	_, err := firestoreClient.Collection(p.collectionName).Doc(id).Delete(ctx)
	if err != nil {
		slog.Error("Error deleting product", "error", err)
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) GetProducts() ([]*model.Product, error) {
	ctx := context.Background()
	firestoreClient := database.GetFirestoreClient()

	// Obtenemos todos los documentos de la colección de productos
	docs, err := firestoreClient.Collection(p.collectionName).Documents(ctx).GetAll()
	if err != nil {
		slog.Error("Error getting products", "error", err)
		return nil, err
	}

	var products []*model.Product
	for _, doc := range docs {
		var product model.Product
		if err := doc.DataTo(&product); err != nil {
			slog.Error("Error mapping product data", "error", err)
			continue
		}
		products = append(products, &product)
	}

	return products, nil
}
