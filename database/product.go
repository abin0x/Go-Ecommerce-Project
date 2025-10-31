package database

// global slice to store products
var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func Store(p Product) {
	productList = append(productList, p)
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product

	for idx, p := range productList {
		if p.ID != productID {
			tempList[idx] = p
		}
	}
	productList = tempList
}

func init() {
	// Initialize the product list with some sample products
	prd1 := Product{
		ID:          1,
		Title:       "Laptop",
		Description: "A high-performance laptop",
		Price:       999.99,
		ImgUrl:      "https://techterms.com/img/xl/laptop_586.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Smartphone",
		Description: "A latest model smartphone",
		Price:       699.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTTYAx3hqK3b3i-JGJlw7nhqqrccJtkMnNv8Q&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Headphones",
		Description: "Noise-cancelling headphones",
		Price:       199.99,
		ImgUrl:      "https://cdn.mos.cms.futurecdn.net/HMCWShKerkfeNQmYYhE3p7.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Smartwatch",
		Description: "A smartwatch with various features",
		Price:       299.99,
		ImgUrl:      "https://istarmax.com/wp-content/uploads/2024/04/Starmax-Product-Range-Summer-2024-2.png",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Tablet",
		Description: "A lightweight tablet",
		Price:       399.99,
		ImgUrl:      "https://p2-ofp.static.pub/fes/cms/2021/10/28/juqs65pgl1gh3dysi7yv1tnvtsiqva364946.png",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Camera",
		Description: "A digital camera",
		Price:       499.99,
		ImgUrl:      "https://www.gearpatrol.com/wp-content/uploads/sites/2/2023/09/best-vintage-film-cameras-refresh-lead-650de2dd54d04-jpg.webp",
	}
	// productList = []Product{prd1, prd2, prd3, prd4, prd5, prd6}
	productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)
}
