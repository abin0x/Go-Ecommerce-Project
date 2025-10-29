class ProductManager {
    constructor() {
        this.apiUrl = 'http://localhost:8080/products';
        this.createApiUrl = 'http://localhost:8080/createproduct';
        this.productsGrid = document.getElementById('productsGrid');
        this.loadingElement = document.getElementById('loading');
        this.errorElement = document.getElementById('error');
        this.addProductForm = document.getElementById('addProductForm');
        this.formMessage = document.getElementById('formMessage');
        this.resetFormBtn = document.getElementById('resetForm');
        
        this.initializeEventListeners();
    }

    initializeEventListeners() {
        // Form submission
        this.addProductForm.addEventListener('submit', (e) => {
            e.preventDefault();
            this.handleFormSubmit();
        });

        // Reset form
        this.resetFormBtn.addEventListener('click', () => {
            this.resetForm();
        });
    }

    async handleFormSubmit() {
        const formData = new FormData(this.addProductForm);
        const productData = {
            title: formData.get('title'),
            description: formData.get('description'),
            price: parseFloat(formData.get('price')),
            imgUrl: formData.get('imgUrl') || ''
        };

        try {
            await this.createProduct(productData);
        } catch (error) {
            this.showFormMessage('Error creating product: ' + error.message, 'error');
        }
    }

    async createProduct(productData) {
        this.showFormMessage('Creating product...', '');

        try {
            const response = await fetch(this.createApiUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(productData)
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const result = await response.json();
            
            this.showFormMessage('Product created successfully!', 'success');
            this.resetForm();
            
            // Refresh the product list
            setTimeout(() => {
                this.fetchProducts();
            }, 1000);

        } catch (error) {
            console.error('Error creating product:', error);
            throw error;
        }
    }

    showFormMessage(message, type) {
        this.formMessage.textContent = message;
        this.formMessage.className = 'form-message';
        if (type) {
            this.formMessage.classList.add(type);
        }
    }

    resetForm() {
        this.addProductForm.reset();
        this.formMessage.textContent = '';
        this.formMessage.className = 'form-message';
    }

    async fetchProducts() {
        try {
            this.showLoading();
            
            const response = await fetch(this.apiUrl);
            
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            
            const products = await response.json();
            this.hideLoading();
            this.displayProducts(products);
            
        } catch (error) {
            console.error('Error fetching products:', error);
            this.showError();
        }
    }

    showLoading() {
        this.loadingElement.style.display = 'block';
        this.errorElement.style.display = 'none';
        this.productsGrid.innerHTML = '';
    }

    hideLoading() {
        this.loadingElement.style.display = 'none';
    }

    showError() {
        this.loadingElement.style.display = 'none';
        this.errorElement.style.display = 'block';
    }

    displayProducts(products) {
        if (!products || !Array.isArray(products) || products.length === 0) {
            this.productsGrid.innerHTML = `
                <div class="no-products" style="grid-column: 1 / -1; text-align: center; padding: 40px;">
                    <h3>No products available</h3>
                    <p>Check back later for new products!</p>
                </div>
            `;
            return;
        }

        this.productsGrid.innerHTML = products.map(product => this.createProductCard(product)).join('');
    }

    createProductCard(product) {
        const productTitle = product?.title || 'Unnamed Product';
        const productPrice = product?.price || 0;
        const productDescription = product?.description;
        const productImage = product?.imgUrl;
        const productId = product?.id;
        
        const imageElement = this.getProductImage(productImage, productTitle);
        
        return `
            <div class="product-card" data-product-id="${productId || ''}">
                <div class="product-image">
                    ${imageElement}
                </div>
                <h3 class="product-title">${this.escapeHtml(productTitle)}</h3>
                <div class="product-price">$${this.formatPrice(productPrice)}</div>
                ${productDescription ? `<p class="product-description">${this.escapeHtml(productDescription)}</p>` : ''}
            </div>
        `;
    }

    getProductImage(imgUrl, title) {
        if (imgUrl && !imgUrl.includes('example.com')) {
            return `<img src="${imgUrl}" alt="${title}" class="product-img" onerror="this.style.display='none'; this.parentNode.innerHTML=window.productManager.getProductIcon('${title}')">`;
        } else {
            return this.getProductIcon(title);
        }
    }

    getProductIcon(title) {
        const icons = {
            laptop: '💻',
            smartphone: '📱',
            phone: '📱',
            headphones: '🎧',
            smartwatch: '⌚',
            watch: '⌚',
            tablet: '📱',
            camera: '📷',
            electronic: '📱',
            computer: '💻',
            default: '📦'
        };

        const searchString = (title || '').toString().toLowerCase();
        
        for (const [key, icon] of Object.entries(icons)) {
            if (searchString.includes(key)) {
                return icon;
            }
        }
        
        return icons.default;
    }

    formatPrice(price) {
        if (typeof price !== 'number') {
            const numPrice = parseFloat(price);
            if (!isNaN(numPrice)) {
                return numPrice.toFixed(2);
            }
            return '0.00';
        }
        return price.toFixed(2);
    }

    escapeHtml(unsafe) {
        if (typeof unsafe !== 'string') {
            if (unsafe === null || unsafe === undefined) return '';
            return String(unsafe);
        }
        return unsafe
            .replace(/&/g, "&amp;")
            .replace(/</g, "&lt;")
            .replace(/>/g, "&gt;")
            .replace(/"/g, "&quot;")
            .replace(/'/g, "&#039;");
    }

    refresh() {
        this.fetchProducts();
    }
}

// Initialize and load products when page loads
document.addEventListener('DOMContentLoaded', () => {
    const productManager = new ProductManager();
    productManager.fetchProducts();
    
    // Make it available globally for debugging
    window.productManager = productManager;
    
    // Add click handler for error message to retry
    document.getElementById('error').addEventListener('click', () => {
        productManager.fetchProducts();
    });
});

// Auto-refresh every 30 seconds
setInterval(() => {
    if (window.productManager) {
        window.productManager.fetchProducts();
    }
}, 30000);