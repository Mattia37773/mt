db = db.getSiblingDB("appDb");

db.createUser({
  user: "appUser",
  pwd: "appPassword",
  roles: [
    {
      role: "readWrite",
      db: "appDb"
    }
  ]
});

// Users
db.users.insertMany([
  {
    _id: ObjectId("64b000000000000000000001"),
    email: "john.doe@test.com",
    name: "John Doe",
    role: "admin",
    active: true,
    createdAt: new Date()
  },
  {
    _id: ObjectId("64b000000000000000000002"),
    email: "jane.smith@test.com",
    name: "Jane Smith",
    role: "user",
    active: true,
    createdAt: new Date()
  },
  {
    _id: ObjectId("64b000000000000000000003"),
    email: "inactive@test.com",
    name: "Inactive User",
    role: "user",
    active: false,
    createdAt: new Date()
  }
]);

// Products
db.products.insertMany([
  {
    _id: ObjectId("64b100000000000000000001"),
    name: "Laptop Pro",
    category: "electronics",
    price: 1299.99,
    stock: 15,
    active: true
  },
  {
    _id: ObjectId("64b100000000000000000002"),
    name: "Mechanical Keyboard",
    category: "accessories",
    price: 129.99,
    stock: 50,
    active: true
  },
  {
    _id: ObjectId("64b100000000000000000003"),
    name: "Old Monitor",
    category: "electronics",
    price: 199.99,
    stock: 0,
    active: false
  }
]);

// Orders
db.orders.insertMany([
  {
    _id: ObjectId("64b200000000000000000001"),
    userId: ObjectId("64b000000000000000000001"),
    status: "completed",
    items: [
      {
        productId: ObjectId("64b100000000000000000001"),
        quantity: 1,
        price: 1299.99
      }
    ],
    total: 1299.99,
    createdAt: new Date()
  },
  {
    _id: ObjectId("64b200000000000000000002"),
    userId: ObjectId("64b000000000000000000002"),
    status: "pending",
    items: [
      {
        productId: ObjectId("64b100000000000000000002"),
        quantity: 2,
        price: 129.99
      }
    ],
    total: 259.98,
    createdAt: new Date()
  }
]);

// Payments
db.payments.insertMany([
  {
    orderId: ObjectId("64b200000000000000000001"),
    provider: "stripe",
    status: "paid",
    amount: 1299.99,
    transactionId: "txn_test_001",
    createdAt: new Date()
  },
  {
    orderId: ObjectId("64b200000000000000000002"),
    provider: "paypal",
    status: "pending",
    amount: 259.98,
    transactionId: "txn_test_002",
    createdAt: new Date()
  }
]);

// Sessions
db.sessions.insertMany([
  {
    userId: ObjectId("64b000000000000000000001"),
    token: "test-token-admin",
    expiresAt: new Date(Date.now() + 86400000)
  },
  {
    userId: ObjectId("64b000000000000000000002"),
    token: "test-token-user",
    expiresAt: new Date(Date.now() + 86400000)
  }
]);

// Notifications
db.notifications.insertMany([
  {
    userId: ObjectId("64b000000000000000000001"),
    type: "order",
    message: "Your order has been completed",
    read: false,
    createdAt: new Date()
  },
  {
    userId: ObjectId("64b000000000000000000002"),
    type: "system",
    message: "Welcome to the platform",
    read: true,
    createdAt: new Date()
  }
]);

// Logs
db.audit_logs.insertMany([
  {
    action: "USER_CREATED",
    entity: "user",
    entityId: ObjectId("64b000000000000000000001"),
    createdAt: new Date()
  },
  {
    action: "ORDER_CREATED",
    entity: "order",
    entityId: ObjectId("64b200000000000000000001"),
    createdAt: new Date()
  }
]);
