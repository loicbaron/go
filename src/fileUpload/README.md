### Challenge: **Build a Simple Distributed File Upload Service**

#### Problem Description:

Create a simple distributed system that simulates a file upload service with multiple servers communicating via HTTP. The architecture should involve:

1. **API Gateway**:

   - Exposes an endpoint (`/upload`) that clients use to upload files.
   - It should handle incoming file uploads and forward the requests to an appropriate **file storage service**.
   - The `API Gateway` should handle retries and error cases in case the **file storage service** is temporarily unavailable.

2. **File Storage Service**:

   - A separate service that simulates storing the files.
   - Each request contains metadata (file name, size, etc.), and the storage service should simulate saving the file.
   - The storage service should expose a simple endpoint (e.g., `/store`) to receive files.

3. **Client**:

   - A client that sends the file to the `API Gateway`.
   - The client should be able to retry the upload if the server is temporarily unavailable.

4. **Error Handling**:
   - Handle network errors (e.g., server down) with retries and appropriate error messages.
   - Include logging at each step of the process to simulate real-world scenarios.

#### Requirements:

1. **Modular Architecture**:

   - Separate the concerns of the `API Gateway` and the `File Storage Service`.
   - Use Go’s `http` package to simulate API communication.
   - Use error handling and retries where necessary.
   - The architecture should be designed with scalability in mind, though the current scope is just a simulation.

2. **API Design**:

   - **API Gateway**:
     - Endpoint: `POST /upload` (file upload)
   - **File Storage Service**:
     - Endpoint: `POST /store` (store file)

3. **Logging**:

   - Implement logging at the `API Gateway` and `File Storage Service` to track request and response data (you can use `log` package).

4. **Testing**:
   - Write tests for the `API Gateway` and `File Storage Service`.
   - Ensure that retry logic works properly when the storage service is temporarily unavailable.

#### Example Flow:

1. **Client** sends a request to `API Gateway` to upload a file.
2. **API Gateway** forwards the file to the **File Storage Service**.
3. If the **File Storage Service** is unavailable, the **API Gateway** should retry the upload (with exponential backoff or a simple retry mechanism).
4. Once the **File Storage Service** receives the file, it responds with a success message, and the **API Gateway** forwards this to the **Client**.

### Example:

#### 1. Client Request:

```bash
POST /upload
Host: api.example.com
Content-Type: multipart/form-data
File: <file_data>
```

#### 2. API Gateway forwards the request to File Storage Service:

```bash
POST /store
Host: storage.example.com
Content-Type: multipart/form-data
File: <file_data>
```

#### 3. File Storage Service responds with success:

```json
{
  "status": "success",
  "message": "File uploaded successfully"
}
```

#### Expected Deliverables:

- **`API Gateway`** implementation in Go with retry logic.
- **`File Storage Service`** implementation in Go that receives and simulates storing files.
- **Client** code to test uploading files.
- **Test Cases** to verify functionality (e.g., retries, error handling).

---

### Bonus:

- Implement **exponential backoff** for retries to avoid overwhelming the storage service when it is temporarily unavailable.
- Implement **health check** endpoints to simulate the file storage service being down and up.

---

This challenge will test your ability to design a **scalable architecture** with proper **error handling** and **API communication** between services. It emphasizes separation of concerns (API Gateway vs. File Storage), retry logic, and testing, which are important aspects of real-world software systems.

Let me know if you'd like more guidance on this challenge!
