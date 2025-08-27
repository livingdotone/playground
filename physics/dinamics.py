import numpy as np
import matplotlib.pyplot as plt
import matplotlib.animation as animation

# --- Block 1: Physics and Simulation Parameters ---
FORCE = 100.0          # Applied force in Newtons (N)
LIGHT_MASS = 10.0      # Mass of the light object in kg
HEAVY_MASS = 50.0      # Mass of the heavy object in kg

DURATION = 10.0        # Total simulation time in seconds
FPS = 25               # Frames per second for the GIF

# --- Physics Calculations ---
# Calculate acceleration for each object using Newton's 2nd Law (a = F/m)
light_acceleration = FORCE / LIGHT_MASS
heavy_acceleration = FORCE / HEAVY_MASS

# --- Block 2: Plotting Setup ---
fig, ax = plt.subplots(figsize=(10, 5))
ax.set_xlim(-10, 210)
ax.set_ylim(0, 3)
ax.set_xlabel("Position (m)")
ax.get_yaxis().set_visible(False)

# Animation objects (the elements that will be updated each frame)
light_block, = ax.plot([], [], 'bo', markersize=12, label=f'Light ({LIGHT_MASS}kg)') 
heavy_block, = ax.plot([], [], 'rs', markersize=12, label=f'Heavy ({HEAVY_MASS}kg)')
info_text = ax.text(0.05, 0.85, '', transform=ax.transAxes, verticalalignment='top')

# Static elements (drawn once)
ax.axhline(1, color='gray', linestyle='--')
ax.axhline(2, color='gray', linestyle='--')
ax.legend(loc='upper right')

# --- Block 3: Animation Functions ---
def init():
    """Initializes the animation by drawing a blank frame."""
    light_block.set_data([], [])
    heavy_block.set_data([], [])
    info_text.set_text('')
    return light_block, heavy_block, info_text

def update(t):
    """
    This function is called for each frame to update the positions 
    of the objects at a given time 't'.
    """
    # Calculate the position for each block using the kinematics formula
    light_position = 0.5 * light_acceleration * t**2
    heavy_position = 0.5 * heavy_acceleration * t**2
    
    # Update the data for each plotted object.
    # The coordinates must be passed as a sequence (e.g., a list with one item).
    light_block.set_data([light_position], [2])
    heavy_block.set_data([heavy_position], [1])
    
    # Update the informational text on the plot
    info_text.set_text(f'Time = {t:.2f} s\n'
                       f'Force = {FORCE} N\n'
                       f'a_light = {light_acceleration:.1f} m/s²\n'
                       f'a_heavy = {heavy_acceleration:.1f} m/s²')
                        
    # Update the plot title for each frame to show the current time
    ax.set_title(f"Newton's 2nd Law (Time: {t:.2f}s)")
    return light_block, heavy_block, info_text

# --- Block 4: Create and Save the Animation ---
# Calculate the total number of frames and the time points for each frame
frame_count = int(DURATION * FPS)
times = np.linspace(0, DURATION, frame_count)

# Create the animation object
ani = animation.FuncAnimation(fig, update, frames=times,
                            init_func=init, blit=False, interval=1000/FPS)

# Attempt to save the animation as a GIF file
try:
    output_filename = 'newtons_law_animation.gif'
    print(f"Attempting to save animation to '{output_filename}'...")
    # The 'pillow' writer is used to create the GIF
    ani.save(output_filename, writer='pillow', fps=FPS)
    print(f"Animation successfully saved as '{output_filename}'!")

except Exception as e:
    print(f"\n--- ERROR ---")
    print(f"An error occurred: {e}")
    print("If it's a 'format not supported' error, please use the PNG frame generation code instead.")