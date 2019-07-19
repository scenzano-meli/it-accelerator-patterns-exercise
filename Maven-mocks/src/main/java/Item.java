public class Item {

    int id;
    String title;
    String category_id;
    float price;
    String currency_id;
    int quantity;
    String condition;
    String [] pictures;


    public Item(int id) {
        this.id = id;
        this.title = "Título del item";
        this.category_id = "MLA" + this.id;
        this.price = 100;
        this.currency_id = "ARS";
        this.quantity = 10;
        this.condition = "New";
        this.pictures = new String[] {"img1.png", "img2.png"};
    }
}



