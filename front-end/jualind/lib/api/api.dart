import 'package:http/http.dart' as http;

class Api {
  Future<String> get(String url, Map<String, String> head) async {
    final response = await http.get(Uri.parse(url), headers: head);
    if (response.statusCode == 200) {
      return response.body;
    } else {
      throw Exception("Gagal");
    }
  }
}
