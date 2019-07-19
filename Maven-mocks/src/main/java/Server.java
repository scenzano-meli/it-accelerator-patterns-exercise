import com.google.gson.Gson;
import org.mockserver.client.server.MockServerClient;
import org.mockserver.model.Delay;
import org.mockserver.model.Header;
import org.mockserver.model.HttpRequest;
import org.mockserver.model.HttpResponse;
import java.util.concurrent.TimeUnit;

import static org.mockserver.integration.ClientAndServer.startClientAndServer;

public class Server {

    static MockServerClient mockServer = startClientAndServer (8081);
    static Gson gson;

    public static void consulta(String metodo, String ruta, int codigo, String content, String body, int delay) {
        mockServer.when (
                HttpRequest.request ().withMethod (metodo)
                                    .withPath (ruta)
        ).respond (
            HttpResponse.response ().withStatusCode (codigo)
                .withHeader (new Header ("Content-Type", content))
                .withBody (body)
                .withDelay (new Delay (TimeUnit.MILLISECONDS, delay))
        );
    }

    public static void main(String[] args) {
        gson = new Gson ();
        Item item = new Item(12345);
        User user = new User();
        Site site = new Site();
        Country country = new Country();

        consulta("GET", "/items/.*", 200, "application/json; charset=utf-8",
                gson.toJson (item), 0);

        consulta ("GET", "/users/.*", 200, "application/json", user.toString (), 0);
        consulta ("GET", "/sites/.*", 200, "application/json", site.toString (), 0);
        consulta ("GET", "/countries/.*", 200, "application/json", country.toString (), 0);

    }
}


