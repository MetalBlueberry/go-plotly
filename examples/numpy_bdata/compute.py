import numpy as np
import json
import base64

# Generate Z data using NumPy
# Let's create a similar 3x5 matrix with some values
z = np.array([
    [1, np.nan, 30, 50, 1],
    [20, 1, 60, 80, 30],
    [30, 60, 1, -10, 20]
])



dtype = z.dtype
bdata = z.tobytes()  # Binary data representation
shape = z.shape

# Encode the binary data as a base64 string
z_bdata_base64 = base64.b64encode(bdata).decode('utf-8')

# Create the dictionary that represents the JSON object
z_data = {
    'dtype': str(dtype),
    'shape': shape,
    'bdata': z_bdata_base64 
}

# Convert the dictionary to a JSON object
json_data = json.dumps(z_data, indent=2)

# Output the JSON object
print(json_data)