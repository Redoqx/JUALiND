import 'dart:convert';

class AllProduct {
  String? msg;
  List<Body>? body;

  AllProduct({this.msg, this.body});

  factory AllProduct.fromMap(Map<String, dynamic> map) {
    return AllProduct(
      msg: map['msg'] ?? '',
      body: List<Body>.from(map['body']?.map((x) => Body.fromMap(x))),
    );
  }

  factory AllProduct.fromJson(String source) =>
      AllProduct.fromMap(json.decode(source));

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['msg'] = this.msg;
    if (this.body != null) {
      data['body'] = this.body!.map((v) => v.toJson()).toList();
    }
    return data;
  }
}

class Body {
  int? iD;
  int? ownerID;
  String? name;
  int? price;
  String? description;
  int? currentQuantity;
  int? quantity;
  List<ImageLoc>? imageLoc;

  Body(
      {this.iD,
      this.ownerID,
      this.name,
      this.price,
      this.description,
      this.currentQuantity,
      this.quantity,
      this.imageLoc});

  factory Body.fromMap(Map<String, dynamic> map) {
    return Body(
      iD: map['ID'] ?? '',
      ownerID: map['OwnerID'] ?? '',
      name: map['Name'] ?? '',
      price: map['Price'] ?? '',
      description: map['Description'] ?? '',
      currentQuantity: map['CurrentQuantity'] ?? '',
      quantity: map['Quantity'] ?? '',
      imageLoc:
          List<ImageLoc>.from(map['ImageLoc']?.map((x) => ImageLoc.fromMap(x))),
    );
  }

  // Body.fromJson(Map<String, dynamic> json) {
  //   iD = json['ID'];
  //   ownerID = json['OwnerID'];
  //   name = json['Name'];
  //   price = json['Price'];
  //   description = json['Description'];
  //   currentQuantity = json['CurrentQuantity'];
  //   quantity = json['Quantity'];
  //   imageLoc = json['ImageLoc'] != null
  //       ? new ImageLoc.fromJson(json['ImageLoc'])
  //       : null;
  // }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['ID'] = this.iD;
    data['OwnerID'] = this.ownerID;
    data['Name'] = this.name;
    data['Price'] = this.price;
    data['Description'] = this.description;
    data['CurrentQuantity'] = this.currentQuantity;
    data['Quantity'] = this.quantity;
    if (this.imageLoc != null) {
      data['ImageLoc'] = this.imageLoc!.map((v) => v.toJson()).toList();
    }
    return data;
  }
}

class ImageLoc {
  String? string;
  bool? valid;

  ImageLoc({this.string, this.valid});

  factory ImageLoc.fromMap(Map<String, dynamic> map) {
    return ImageLoc(
      string: map['String'] ?? '',
      valid: map['Valid'] ?? '',
    );
  }

  ImageLoc.fromJson(Map<String, dynamic> json) {
    string = json['String'];
    valid = json['Valid'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['String'] = this.string;
    data['Valid'] = this.valid;
    return data;
  }
}
