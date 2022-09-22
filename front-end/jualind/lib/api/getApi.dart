import 'api.dart';

import 'package:jualind/modal/allProduct.dart';

class GetApi {
  final Api? api = Api();

  Future<AllProduct> getAllProduct() async {
    final url = 'localhost:8000/api/v1/login';
    Map<String, String> head = {
      'Authorization':
          'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklEIjoxLCJOYW1lIjoiQXJpZWYiLCJQYXNzd29yZCI6IiQyYSQwOSRscDhrZFQwZndKcFpVdVVrRkJqZmxPQy5ON2RqbUduZ2JXTnpoSFZlejB2NVNFTkk5N3I4cSIsIkVtYWlsIjoiY29jb0BnbWFpbC5jb20iLCJSb2xlIjoicGVuanVhbCIsIkltYWdlTG9jIjp7IlN0cmluZyI6IiIsIlZhbGlkIjpmYWxzZX19LCJleHAiOjE2NjM3ODkwNjl9.jdG4PpNcSwnQOXagtRDjwH0JYa5Kh7uA-cDNIGZLUds'
    };
    final response = await api!.get(url, head);
    final allProduct = AllProduct.fromJson(response);

    return allProduct;
  }
}
