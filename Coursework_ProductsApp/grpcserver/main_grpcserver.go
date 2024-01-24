package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"main/configuration"
	"main/dbstore"
	gc "main/grpcconnection"
	"main/models"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var db dbstore.DbStore

type server struct {
	gc.UnimplementedProductsServer
}

func mapAttributes(attrs []models.Attribute) []*gc.Attribute {
	output := make([]*gc.Attribute, 0)
	for _, attr := range attrs {
		grpcAttr := gc.Attribute{
			SkuId:     attr.SkuId,
			Key:       attr.Key,
			Value:     attr.Value,
			ValueType: attr.ValueType,
		}
		output = append(output, &grpcAttr)
	}
	return output
}

func mapSkus(skus []models.Sku) []*gc.Sku {
	output := make([]*gc.Sku, 0)

	for _, sku := range skus {
		grpcSku := gc.Sku{
			Id:        sku.Id,
			ProductId: sku.ProductId,
			Amount:    sku.Amount,
			Price:     sku.Price,
			Unit:      sku.Unit,
			Created:   timestamppb.New(sku.Created),
		}
		if sku.Updated.Valid {
			grpcSku.Updated = timestamppb.New(sku.Updated.Time)
		}
		if sku.Archived.Valid {
			grpcSku.Updated = timestamppb.New(sku.Archived.Time)
		}
		grpcSku.Attributes = mapAttributes(sku.Attributes)
		output = append(output, &grpcSku)
	}

	return output
}

func mapProducts(products []models.Product) []*gc.Product {
	output := make([]*gc.Product, 0)
	for _, p := range products {
		grpcP := gc.Product{
			Id:          p.Id,
			CatalogId:   p.CatalogId,
			Name:        p.Name,
			Description: p.Description,
			Created:     timestamppb.New(p.Created),
		}
		if p.Updated.Valid {
			grpcP.Updated = timestamppb.New(p.Updated.Time)
		}
		if p.Archived.Valid {
			grpcP.Archived = timestamppb.New(p.Archived.Time)
		}
		grpcP.Skus = mapSkus(p.Skus)
		output = append(output, &grpcP)
	}
	return output
}

func (s *server) GetProductsByCatalogIdWithFilter(ctx context.Context, req *gc.Request) (*gc.Response, error) {
	log.Printf("Received: %v", req)
	products, err := db.Products.GetProductsByCatalogIdWithFilter(req.CatalogId, req.Filter.Properties)
	if err != nil {
		return nil, err
	}
	return &gc.Response{Products: mapProducts(products)}, nil
}

func main() {
	fmt.Println("Yoo GRPC Server")

	configuration.Setup()
	db = dbstore.GetNewDbStore(configuration.ConnectionString)
	db.Ping()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configuration.GrpcServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	gc.RegisterProductsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
